
<script>
    import {onMount} from 'svelte';
    import {createEventDispatcher} from 'svelte';
    const dispacher = createEventDispatcher();

    export let id = "myarea";
    let hover = false;
    let input = null;
    let form = null;


    onMount(() => {
        input = document.getElementById(id);
        form = input.parentNode;
    })
    
    async function handleSubmit(event){

        let response = await fetch('http://localhost:7070/images', {
            method: 'POST',
            body: new FormData(form)
        });

        if(response.ok){
            dispacher('success')
        }
    }

    function handleChange(event){
        form.dispatchEvent(new Event('submit'));
    }

    function handleDrop(){
        input.files = event.dataTransfer.files;
        input.dispatchEvent(new Event('change'));
        hover = false;
    }
</script>




<form 
    class="drop-area_form"
    enctype="multipart/form-data"
    on:submit|preventDefault={handleSubmit}
    method="POST" 
    on:click={(event)=>{
        event.stopPropagation();
        
    }}>
    <input 
        id="{id}" 
        class="drop-area_input" 
        type="file" name="file" 
        on:change|preventDefault={handleChange}/>
    <label 
        for="{id}" class:hover={'hover'} 
        class="drop-area" on:dragenter|preventDefault={()=>hover=true} 
        on:dragover|preventDefault={()=>hover=true} 
        on:dragleave|preventDefault={()=>hover=false} 
        on:drop|preventDefault={handleDrop}>
        <div class="drop-area_info">
            <slot>
                <span>Arrastre y suelte archivos aquí, o de clic</span>
            </slot>
        </div>
    </label>
</form>

<style>
    .drop-area_form{
        width: 100%;
        display: flex;
    }

    .drop-area{
        width: 100%;
        height: 100%;
        cursor: pointer;
        min-height: 200px;
        border: 2px dashed #ccc;
    }

    .drop-area_input{
        display: none;
    }

    .drop-area_info{
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100%;
    }


</style>