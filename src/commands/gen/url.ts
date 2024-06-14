import clipboard from "clipboardy";
import { Argument, Option } from "commander";
import pc from "picocolors";
import { CommandBase } from "../base.js";
import { CommonGenOptions, configOption, variantsOption } from "../generate.js";
import { API_URL, FONT_URL } from "@/urls.js";
import { checkVariantsString, fetchWithKey, isErrorResponse, toTitleCase } from "@/utils.js";
import { ConfigProvider } from "./configProvider.js";

interface UrlCommandOptions extends CommonGenOptions {
  copy: boolean;
  cssImport: boolean;
}

export class UrlCommand extends CommandBase {
  readonly url = new URL(API_URL);
  readonly cfgProvider = new ConfigProvider();

  async handle(fontName: string, { copy, cssImport, variants, config }: UrlCommandOptions) {
    const fontUrl = new URL(FONT_URL);

    if (!config && !fontName)
      this.error("error: Font family is required");

    let _config: Config | null = null;
    
    if (config)
      try {
        _config = await this.cfgProvider.useFile(config);
      } catch (e) {
        this.error("error: " + (e as Error).message);
      }

    if (_config) {
      copy = _config.copy || false;
      cssImport = _config.toCssImport || false;
    }

    const families: string[] = [];

    const optionItems: OptionItem[] = config && _config ? _config.options : [
      { 
        fontFamily: fontName,
        get variants() {
          return (variants ? variants.split(",") : ["regular"]) as Variant[];
        }
      }
    ];

    for (const { fontFamily, variants } of optionItems) {
      const url = this.url;
      url.searchParams.set("family", toTitleCase(fontFamily));
      
      const v = variants.join(",");

      if (variants)
        checkVariantsString(v);

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

      let hasNonExistentVariant = false;

      if (v !== "all")
        for (const v of variants) {
          if (!Object.hasOwn(font.files, v.endsWith("i") ? v.replace("i", "italic") : v)) {
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

      families.push(font.family.replaceAll(" ", "+") + ":" + (() => {
        if (!config)
          if (v === 'all')
            return Object.keys(font.files).join(",");

        return (v ? this.removeDups(v) : "regular");
      })());
    }

    fontUrl.searchParams.set("display", "swap");
    fontUrl.searchParams.set("family", families.join("|"));

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
      configOption.conflicts(["copy", "cssImport", "variants"]),
      new Option("--copy", "Copy URL to clipboard"),
      new Option("-i, --css-import", "Display as CSS import rule"),
      variantsOption
    ];

    super("url");

    for (const opt of options)
      this.addOption(opt);

    this
      .addArgument(
        new Argument("<font name>", "Font family to generate").argOptional()
      )
      .description("Generate URL for CSS usage");
  }
}
