type TwMergeInput = string | string[] | { [key: string]: boolean };

/**
 * Merge tailwindcss classes into a single string
 */
export function twMerge(...classes: TwMergeInput[]): string {
	return classes
		.map((c) => {
			if (typeof c === 'string') {
				return c;
			}

			if (Array.isArray(c)) {
				return twMerge(...c);
			}

			return Object.keys(c)
				.filter((key) => c[key])
				.join(' ');
		})
		.join(' ');
}
