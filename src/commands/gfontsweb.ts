import open from "open";
import ora from "ora";
import pc from "picocolors";
import { CommandBase } from "./base.js";

export class GfontswebCommand extends CommandBase {
  readonly gfontsUrl = "https://fonts.google.com";

  constructor() {
    super("gfontsweb");
    this.description("Open Google Fonts website");
  }

  async handle() {
    const spinner = ora("Opening website...").start();
    await open(this.gfontsUrl, { wait: true });
    spinner.stop();
    console.log(pc.bold(pc.green("Opened")));
  }
}
