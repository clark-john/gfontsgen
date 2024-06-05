import { Option } from "commander";
import { writeFile } from "fs/promises";
import ora from "ora";
import pc from "picocolors";
import { CommandBase } from "./base.js";
import { API_URL, SPECIMEN_URL } from "@/urls.js";
import { fetchWithKey } from "@/utils.js";

interface ListCommandOptions {
  classification: string;
  decorativeStroke: string;
  limit: number;
  search: string;
  writeToFile: string;
}

export class ListCommand extends CommandBase {
  constructor() {
    const options = [
      new Option("-c, --classification <classification>", "Classification")
        .choices(["display", "handwriting", "mono", "monospace"]),
      new Option("-d, --decorative-stroke <decorative-stroke>", "Classification")
        .choices(["serif", "sans-serif", "slab-serif"]),
      new Option("-l, --limit <number>", "Limit number of search items")
        .argParser(parseInt),
      new Option("-s, --search <keyword>", "Search for a font"),
      new Option("--write-to-file <file>", "Write output to file")
    ];

    super("list");

    for (const option of options)
      this.addOption(option);

    this.description("Get a list of fonts from Google Fonts");
  }

  public async handle({ 
    limit, search, writeToFile, classification, decorativeStroke 
  }: ListCommandOptions) {
    const loading = ora("Loading list of fonts...").start();

    const resp = await fetchWithKey(new URL(API_URL));
    const fullResp = await resp.json() as FullResponse;

    loading.stop();

    let items = [...fullResp.items];

    if (search)
      items = items.filter(val => val.family.toLowerCase().includes(search.toLowerCase()));

    if (decorativeStroke || classification) {
      if (classification === "mono")
        classification = "monospace";

      items = items.filter(val => 
        val.category === decorativeStroke ||
        val.category === classification
      );
    }

    if (limit)
      items = items.slice(0, limit);

    const lines: string[] = [];

    let num = 1;

    if (!items.length) {
      console.error("No results");
      return;
    }

    for (const fontItem of items) {
      lines.push("#" + num);
      lines.push("Name: " + fontItem.family);
      lines.push("Variants: " + Object.keys(fontItem.files).join(", "));

      lines.push(
        "Preview URL: " 
        + SPECIMEN_URL  
        + fontItem.family.replaceAll(" ", "+")
      );

      lines.push("");

      num++;
    }

    const outputData = lines.join("\n");

    if (writeToFile) {
      await writeFile(writeToFile, outputData, "utf8");
      console.log(pc.green("Output written to " + writeToFile));
      return;
    }

    console.log(outputData);
  }
}
