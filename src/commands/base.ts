import { Command } from "commander";

export abstract class CommandBase extends Command {
  abstract handle(...args: any[]): Promise<void>;
}
