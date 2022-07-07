<script>
	import { writable } from 'svelte/store'
	import Todo from "./Todo.svelte";
	let title = ''
	let todo_list = writable([])
	let id = 0

	function CreateToDo (todo) {
		if (!title.trim()) {
			return
		}
		$todo_list.push({
			id,
			title
		})
		$todo_list = $todo_list
		title = ''
		id += 1
	}
</script>

<input type="text"
	   bind:value={title}
	   on:keydown={(e) => {e.key === 'Enter' && CreateToDo()}}
/>
<button on:click={CreateToDo}>
	CreateToDo
</button>

{#each $todo_list as todo}
	<Todo {todo_list} {todo}/>
{/each}