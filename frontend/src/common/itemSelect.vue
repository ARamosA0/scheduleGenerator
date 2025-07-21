<template>
    <DataView :value="props.data">
        <template #list="slotProps">
            <div class="flex flex-col">
                <div v-for="(item, index) in slotProps.items" :key="index">
                    <div
                        class="flex flex-col sm:flex-row sm:items-center p-6 gap-4"
                    >
                        <slot
                            name="item"
                            :item="item"
                            :onToggle="() => $emit('toggle', item)"
                        />
                    </div>
                </div>
            </div>
        </template>
    </DataView>
</template>
<script setup lang="ts">
import { DataView } from "primevue";
const props = defineProps({
    data: {
        type: Array,
        required: true,
    },
});
defineEmits(["toggle"]);
</script>
