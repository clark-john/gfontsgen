import { createWriteStream, promises as fsPromises } from "fs";
import path from "path";
import pc from "picocolors";
import { Readable, promises } from "stream";
import { capitalize, exists, toTitleCase } from "./utils.js";

const { constants: { O_CREAT }, mkdir, open, writeFile } = fsPromises;

export async function fetch2(url: string | URL | Request, requestInit?: RequestInit): Promise<Response> {
  try {
    return await fetch(url, requestInit);
  } catch (e) {
    const err = e as TypeError;
    await writeFile("error.log", err.cause + "\n" + err.stack!);

    console.error("An error occurred while processing the request!");
    process.exit(-1);
  }
}

export async function download(
  fontName: string, variant: string, _path: string, link: string
) {
  const f = toTitleCase(fontName).replaceAll(" ", "");

  const startsWithNumber = /^\d+/.test(variant);
  const isItalic = variant.endsWith("italic");
  
  const v = !startsWithNumber
    ? capitalize(variant)
    : variant.substring(0, 3) + (isItalic ? "Italic" : "Regular");

  const resp = await fetch2(link);

  if (!resp.ok) {
    console.error(pc.bold(pc.red("Failed to download")));
    return;
  }

  const filename = path.resolve(_path, f + "-" + v + path.extname(link));

  if (!(await exists(_path)))
    await mkdir(_path);

  const file = await open(filename, O_CREAT);

  const fileStream = createWriteStream(
    filename, 
    { autoClose: true, encoding: "binary" }
  );

  await promises.pipeline(Readable.fromWeb(resp.body!), fileStream);
  file.close();

  console.log(pc.bold(pc.green("Successfully downloaded " + path.basename(filename))));
}
