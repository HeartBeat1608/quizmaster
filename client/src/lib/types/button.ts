export type ButtonStyle = 'default' | 'action' | 'danger' | 'link' | 'success';

export const baseButtonStyle =
	'px-4 py-2 rounded-md transition-all disabled:opacity-50 disabled:cursor-not-allowed ';

export const buttonStyles: { [key in ButtonStyle]: string } = {
	action: baseButtonStyle + 'bg-indigo-500 text-white hover:bg-indigo-600',
	default: baseButtonStyle + 'bg-gray-200 text-gray-800 hover:bg-gray-300',
	danger: baseButtonStyle + 'bg-red-500 text-white hover:bg-red-600',
	link: baseButtonStyle + 'bg-transparent text-gray-800 hover:bg-gray-200',
	success: baseButtonStyle + 'bg-green-500 text-white hover:bg-green-600'
};
