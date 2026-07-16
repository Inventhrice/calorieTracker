<script lang="ts">

import { defineComponent } from 'vue';

export default defineComponent({
    props: {
        message: { type: String, default: "Generic Error Message." },
        showErr: { type: Boolean, default: false }
    },
    data() {
        return {
            queue: [] as string[]
        }
    },
    methods: {
        async startDisplay(message: string, timeout: number = 5000) {
            this.queue.push(message)
            await setTimeout(() => { 
                const i = this.queue.indexOf(message)
                this.queue.splice(i, 1) 
            }, timeout)
            this.$emit('added')
        }
    },
    created() {
        this.$watch('showErr', () => {
            if (this.showErr) {
                this.startDisplay(this.message)
            }
        })
    },
    emits: {
        added: null
    }

})

</script>

<template>
    <div class="absolute h-full self-center flex flex-col-reverse">
        <TransitionGroup>
            <div v-for="message in queue" :key="message" class="w-fit p-3 rounded-2xl bg-red-500 text">
                <span class="icon mdi--error align-middle scale-125 mr-2"></span>
                <span class="align-middle">{{ message }} Check console for details.</span>
            </div>
        </TransitionGroup>
    </div>
</template>

<style>
.v-enter-active,
.v-leave-active {
    transition: all 0.5s ease;
}

.v-enter-from,
.v-leave-to {
    opacity: 0;
    transform: translateY(30px);
}
</style>