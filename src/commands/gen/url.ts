import clipboard from "clipboardy";
import { Option } from "commander";
import pc from "picocolors";
import { CommandBase } from "../base.js";
import { CommonGenOptions, configOption, variantsOption } from "../generate.js";
import { API_URL, FONT_URL } from "@/urls.js";
import { checkVariantsString, fetchWithKey, isErrorResponse, toTitleCase } from "@/utils.js";

interface UrlCommandOptions extends CommonGenOptions {
  copy: boolean;
  cssImport: boolean;
}

export class UrlCommand extends CommandBase {
  async handle(fontName: string, { copy, cssImport, variants, config }: UrlCommandOptions) {
    const url = new URL(API_URL);
    url.searchParams.set("family", toTitleCase(fontName));

    if (variants)
      checkVariantsString(variants);

    const resp = await fetchWithKey(url);
    const json = await resp.json();

    if (isErrorResponse(json)) {
      const err = (json as ErrorResponse).error;
      if (err.code === 404) {
        console.error("Font family not found");
        return;
      }
    }

    const font = (json as FullResponse).items[0];

    const fontUrl = new URL(FONT_URL);

    let hasNonExistentVariant = false;

    if (variants !== "all")
      for (const v of variants.split(",")) {
        if (!Object.hasOwn(font.files, v)) {
          console.warn(
            pc.bold(pc.yellow("Warning: variant %s doesn't exist on font %s.")),
            v, font.family
          );
          if (!hasNonExistentVariant)
            hasNonExistentVariant = true;
        }
      }

    if (hasNonExistentVariant)
      return;

    const family = font.family.replaceAll(" ", "+") + ":" + (() => {
      if (variants === 'all')
        return Object.keys(font.files).join(",");

      return (variants ? this.removeDups(variants) : "regular");
    })();

    fontUrl.searchParams.set("display", "swap");
    fontUrl.searchParams.set("family", family);

    fontUrl.search = decodeURIComponent(fontUrl.search);

    let final = fontUrl.href;

    if (cssImport)
      final = `@import url("${final}");`;

    console.log(final);

    if (copy) {
      await clipboard.write(final);
      console.log(pc.bold(pc.green("Successfully copied!")));
    }
  }

  private removeDups(variants: string): string {
    return [...new Set(variants.split(",")).values()].join(",");
  }

  constructor() {
    const options = [
      configOption,
      new Option("--copy", "Copy URL to clipboard"),
      new Option("-i, --css-import", "Display as CSS import rule"),
      variantsOption
    ];

    super("url");

    for (const opt of options)
      this.addOption(opt);

    this
      .argument("<font name>", "Font family to generate")
      .description("Generate URL for CSS usage");
  }
}
