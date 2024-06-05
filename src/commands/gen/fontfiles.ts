import { Option } from "commander";
import { createWriteStream } from "fs";
import { access, constants, mkdir, open } from "fs/promises";
import path from "path";
import pc from "picocolors";
import { Readable } from "stream";
import { pipeline } from "stream/promises";
import { CommandBase } from "../base.js";
import { CommonGenOptions, configOption, variantsOption } from "../generate.js";
import { API_URL } from "@/urls.js";
import { 
  capitalize, 
  checkVariantsString, 
  fetchWithKey,
  isErrorResponse, 
  toTitleCase 
} from "@/utils.js";

interface FontFilesCommandOptions extends CommonGenOptions {
  path: string;
  woff: boolean;
}

export class FontFilesCommand extends CommandBase {
  async handle(fontName: string, { path, variants, woff }: FontFilesCommandOptions) {
    const url = new URL(API_URL);
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
        await this.download(fontName, variant, path, link);
      return;
    }

    for (const variant of variants ? variants.split(",") : ["regular"]) {
      const v = variant.endsWith("i") ? variant.replace("i", "italic") : variant;
      
      if (!Object.hasOwn(font.files, v)) {
        console.log(pc.bold(pc.yellow("Warning: variant %s doesn't exist on font %s")), v, font.family);
        continue;
      };

      await this.download(fontName, v, path, font.files[v]);
    }
  }

  async download(fontName: string, variant: string, _path: string, link: string) {
    const f = capitalize(fontName.replaceAll(" ", ""));

    const startsWithNumber = /^\d+/.test(variant);
    const isItalic = variant.endsWith("italic");
    
    const v = !startsWithNumber
      ? capitalize(variant)
      : variant.substring(0, 3) + (isItalic ? "Italic" : "Regular");

    const resp = await fetch(link);

    if (!resp.ok) {
      console.error(pc.bold(pc.red("Failed to download")));
      return;
    }

    const filename = path.resolve(_path, f + "-" + v + path.extname(link));

    try {
      await access(_path);
    } catch (e) {
      mkdir(_path);
    }

    const file = await open(filename, constants.O_CREAT);

    const fileStream = createWriteStream(
      filename, 
      { autoClose: true, encoding: "binary" }
    );

    await pipeline(Readable.fromWeb(resp.body!), fileStream);
    file.close();
  }
  
  constructor() {
    const options = [
      configOption,
      new Option("-p, --path <folder>", "Output path to generate fonts").default("fonts"),
      variantsOption,
      new Option("--woff", "use WOFF2 format")
    ];

    super("fontfiles");

    for (const opt of options)
      this.addOption(opt);

    this
      .argument("<font name>", "Font family to generate")
      .description("Generate font files");
  }
}
