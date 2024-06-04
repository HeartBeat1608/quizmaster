<script lang="ts">
	import type { ButtonStyle } from '$lib/types/button';
	import { buttonStyles } from '$lib/types/button';
	import { twMerge } from '$lib/utils/twMerge';
	import { createEventDispatcher } from 'svelte';

	export let variant: ButtonStyle = 'default';

	$: buttonStyle = buttonStyles[variant];

	const dispatcher = createEventDispatcher();

	function dispatch(eventName: string) {
		return (event: any) => dispatcher(eventName, event);
	}
</script>

<button {...$$props} class={twMerge(buttonStyle, $$props.class || '')} on:click={dispatch('click')}>
	<slot></slot>
</button>
