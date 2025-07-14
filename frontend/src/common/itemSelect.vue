<template>
    <DataView :value="props.data">
        <template #list="slotProps">
            <div class="flex flex-col">
                <div v-for="(item, index) in slotProps.items" :key="index">
                    <div
                        class="flex flex-col sm:flex-row sm:items-center p-6 gap-4"
                    >
                        <Checkbox v-model="item.status" binary />
                        <div class="flex flex-col">
                            <p class="text-lg font-bold">
                                {{ item.title }}
                            </p>
                            <p>{{ item.subtitle }}</p>
                            <p class="text-sm">
                                {{ item.detalle }}
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </template>
    </DataView>
</template>
<script setup lang="ts">
import { DataView, Checkbox } from "primevue";
const props = defineProps({
    data: {
        type: Array,
        required: true,
    },
});
</script>
