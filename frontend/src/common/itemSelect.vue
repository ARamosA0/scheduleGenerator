<template>
    <input
        v-model="search"
        type="text"
        placeholder="Buscar..."
        class="border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
    />
    <DataView :value="filteredData">
        <template #list="slotProps">
            <div class="flex flex-col max-h-[300px] overflow-y-auto">
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
import { ref, computed } from "vue";
const props = defineProps({
    data: {
        type: Array,
        required: true,
    },
});
defineEmits(["toggle"]);

const search = ref();
const filteredData = computed(() => {
    if (search.value === null || search.value === undefined) return props.data;
    if (!search.value.trim()) return props.data;
    return props.data.filter((item) =>
        JSON.stringify(item).toLowerCase().includes(search.value.toLowerCase()),
    );
});
</script>
