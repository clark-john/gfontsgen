import { Argument, Option } from "commander";
import { rm } from "fs/promises";
import pc from "picocolors";
import { CommandBase } from "../base.js";
import { CommonGenOptions, configOption, variantsOption } from "../generate.js";
import { download } from "@/http.js";
import { API_URL } from "@/urls.js";
import { 
  checkVariantsString, 
  exists, 
  fetchWithKey,
  isErrorResponse, 
  toTitleCase 
} from "@/utils.js";
import { ConfigProvider } from "./configProvider.js";

interface FontFilesCommandOptions extends CommonGenOptions {
  path: string;
  woff: boolean;
}

export class FontFilesCommand extends CommandBase {
  readonly url = new URL(API_URL);
  readonly cfgProvider = new ConfigProvider();

  async handle(fontName: string, { path, variants, woff, config }: FontFilesCommandOptions) {
    const cfg = this.cfgProvider;
    
    if (!config && !fontName)
      this.error("error: Font family is required");

    if (!config)
      return this.processFont(fontName, { path, variants, woff });

    if (config) {
      let _config: Config;
      try {
        _config = await cfg.useFile(config);
      } catch (e) {
        this.error("error: " + (e as Error).message);
      }

      if (_config) {
        if (_config.outputPath)
          path = _config.outputPath;

        if (_config.deleteFontDir)
          if (await exists(path))
            await rm(path, { recursive: true, force: true });
      }

      for (const { fontFamily, variants: v } of _config.options)
        this.processFont(fontFamily, { path, variants: v.join(","), woff });
    }
  }

  async processFont(
    fontName: string, 
    { path, variants, woff }: Omit<FontFilesCommandOptions, "config">,
  ) {
    
    const url = this.url;
    url.searchParams.set("family", toTitleCase(fontName));
    
    if (woff)
      url.searchParams.set("capability", "woff2");

    /* start of repeated code */
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
    /* end of repeated code */

    if (variants === 'all') {
      for (const [variant, link] of Object.entries(font.files))
        await download(fontName, variant, path, link);
      return;
    }

    for (const variant of variants ? variants.split(",") : ["regular"]) {
      const v = variant.endsWith("i") ? variant.replace("i", "italic") : variant;
      
      if (!Object.hasOwn(font.files, v)) {
        console.log(pc.bold(pc.yellow("Warning: variant %s doesn't exist on font %s")), v, font.family);
        continue;
      };

      await download(fontName, v, path, font.files[v]);
    }
  }
  
  constructor() {
    const options = [
      configOption.conflicts(["woff", "path", "variants"]),
      new Option("-p, --path <folder>", "Output path to generate fonts").default("fonts"),
      variantsOption,
      new Option("--woff", "use WOFF2 format")
    ];

    super("fontfiles");

    for (const opt of options)
      this.addOption(opt);

    this
      .addArgument(new Argument("<font name>", "Font family to generate").argOptional())
      .description("Generate font files");
  }
}
