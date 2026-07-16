<script>
import { api_call, api_get } from "../../js/api.js"
import { calc_day_goals } from "./goals.ts";
import User_info from "./user_info.vue";
import Goals from "./goals.vue";

export default {
    components: { Goals, User_info },
    data() {
        return {
            title: "Settings"
        }
    }
}
</script>

<template>
    <div class="content-list items-start space-y-3">
        <div class="module-background containerDefaults flex flex-col"><user_info></user_info></div>
        <div class="module-background containerDefaults"><goals></goals></div>        
    </div>
</template>

<style>
@import "tailwindcss";

.informationSpan {
    @apply flex mx-2 place-items-center;
}

.containerDefaults {
    @apply rounded-xl w-fit h-fit p-2;
}
</style>
