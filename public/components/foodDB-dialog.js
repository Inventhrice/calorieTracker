export default {
    props: {
        selected: Object
    },
    template: `
        <dialog class="dialog text flex flex-col" open>
                <label for="name">Name</label>
                <input type="text" class="title-text mx-0 py-2 dialog-input" id="name" v-model="selected.name" />

                <div class="grid grid-cols-2 gap-3">
                    <label for="calPerG">Calorie (g)</label>
                    <label for="proteinPerG">Protein (g)</label>
                    <input type="number" class="dialog-input" step="0.01" v-model="selected.calperg" />
                    <input type="number" class="dialog-input" step="0.01" id="proteinPerG" v-model="selected.proteinperg"/>
                    <label for="fatPerG">Fat (g)</label>
                    <label for="carbPerG">Carbs (g)</label>
                    <input type="number" class="dialog-input" step="0.01" id="fatPerG" v-model="selected.fatperg"/>
                    <input type="number" class="dialog-input" step="0.01" id="carbPerG" v-model="selected.carbperg"/>    
                </div>                
                
                <label for="notes">Notes</label>
                <textarea id="notes" class="dialog-input" v-model="selected.notes"></textarea>
            
                <label for="source">Source</label>
                <textarea id="source" class="dialog-input" v-model="selected.source" />

                <div class="flex justify-end">
                    <button class="btn" @click="$emit('close-dialog')">Cancel</button>
                    <button class="btn btn-confirm" @click="$emit('confirm-dialog')">Save</button>
                </div>
        </dialog>
    `
}