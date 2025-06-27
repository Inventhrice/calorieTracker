export default {
    template: `
        <dialog class="dialog text" open>
            Are you sure you want to perform this action?
            <div class="flex justify-end">
                <button class="btn" @click="$emit('close-dialog')">Cancel</button>
                <button class="btn btn-uhoh" @click="$emit('confirm-dialog')">Delete</button>
            </div>
        </dialog>
    `
}