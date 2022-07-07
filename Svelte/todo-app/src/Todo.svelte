<script>
    export let todo_list
    export let todo // props

    let isEdit = false
    let title = ""

    function onEdit() {
        isEdit = true
        title = todo.title
    }

    function editOff() {
        isEdit = false
    }

    function updateTodo() {
        isEdit = false 
        todo.title = title
    }

    function deleteTodo() {
        $todo_list = $todo_list.filter(t => t.id !== todo.id)
    }
</script>

{#if isEdit}
    <div>
        <input bind:value={title}
               type="text" 
        on:keydown={(e) => {e.key === 'Enter' && CreateToDo()}}
        />
        <button on:click={updateTodo}>
            OK
        </button>
        <button on:click={editOff}>
            Cancel
        </button>
    </div>

{:else}
<div>
    {todo.title}
    <button on:click={onEdit}>
        Edit
    </button>
    <button on:click={deleteTodo}>
        Delete
    </button>
</div>
{/if}