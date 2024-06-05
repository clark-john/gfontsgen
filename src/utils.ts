import { EOL } from "os";
import pc from "picocolors";

export function toTitleCase(text: string): string {
	return text
		.split(" ")
		.map(capitalize)
		.join(" ");
}

export function capitalize(text: string): string {
	if (!text)
		return "";
	return text[0].toUpperCase() + text.substring(1);
}

export function fetchWithKey(url: URL, requestInit?: RequestInit) {
	url.searchParams.set("key", process.env.GFONTSGEN_API_KEY!);
	return fetch(url, requestInit);
}

export function isErrorResponse(json: any): boolean {
	return Object.hasOwn(json as Record<string, any>, "error");
}

const variantsGroupRegex = /^[\w,]*$/;
const variantRegex = /^([1-9]00[iI]?(?!.))|regular|italic|all$/;

function write(text: string) {
	process.stdout.write(text);
}

export function checkVariantsString(variants: string): string[] {
	if (!variantsGroupRegex.test(variants)) {
		console.error(pc.red(`Variants argument must be in this format: "400,500,600"`));
		process.exit(-1);
	}

	const vars = variants.split(",");

	const indices: number[] = [];

	vars
		.forEach((variant, index) => {
			if (!variantRegex.test(variant))
				indices.push(index);
		});

	if (indices.length) {
		write("Following variants are invalid: [");
		
		let i = 0;
		for (const index of indices) {
			write(vars[index]);
			i++;
			if (i < indices.length) {
				write(",");
			}
		}
		write("]" + EOL);
	
		process.exit(-1);
	}

	return vars;
}
