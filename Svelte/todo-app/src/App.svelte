<script>
	import Todo from "./Todo.svelte";
	let title = ''
	let todo_list = []
	let id = 0

	function CreateToDo (todo) {
		todo_list.push({
			id,
			title
		})
		todo_list = todo_list
		title = ''
		id += 1
	}
</script>

<input type="text"
	   bind:value={title} />
<button on:click={CreateToDo}>
	CreateToDo
</button>

{#each todo_list as todo}
	<Todo {todo} />
{/each}