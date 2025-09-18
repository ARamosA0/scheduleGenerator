<template>
    {{ props.items }}
    <div id="app">
        <CalendarView
            :show-date="data.showDate"
            :items="props.items"
            :starting-day-of-week="data.startingDayOfWeek"
            :display-period-uom="props.displayPeriodUom"
            class="theme-default"
        >
            <template #header="{ headerProps }">
                <CalendarViewHeader
                    slot="header"
                    :header-props="headerProps"
                    @input="setShowDate"
                />
            </template>
        </CalendarView>
    </div>
</template>
<script setup lang="js">
import { ref, watch } from "vue";
import { CalendarView, CalendarViewHeader } from "vue-simple-calendar";

import "../../node_modules/vue-simple-calendar/dist/vue-simple-calendar.css";
import "../../node_modules/vue-simple-calendar/dist/css/default.css";

const props = defineProps({
    displayPeriodUom: {
        type: Object,
        deault: null,
    },
    items: {
        type: Object,
        default: null,
    },
});

watch(
    () => props.items,
    (newValue) => {
        if (newValue !== null) {
            const transformed = newValue.map((item, key) => {
                console.log("item", item);
                return {
                    ...item,
                    startDate: fromDateString(item.startDate),
                    endDate: fromDateString(item.endDate),
                };
            });
            data.value.items = transformed;
        }
    },
    { immediate: true },
);

const fromDateString = (dateStr) => {
    console.log("dataStr", dateStr);
    if (dateStr.includes(" ")) {
        const [datePart, timePart] = dateStr.split(" ");
        const [year, month, day] = datePart.split("-").map(Number);
        const [hour, minute] = timePart.split(":").map(Number);

        return new Date(year, month - 1, day, hour, minute);
    } else {
        const [year, month, day] = dateStr.split("-").map(Number);
        return new Date(year, month - 1, day, 0, 0); // Hora 00:00
    }
};

const thisMonth = (d, h, m) => {
    const t = new Date();
    return new Date(t.getFullYear(), t.getMonth(), d, h ?? 0, m ?? 0);
};

const data = ref({
    showDate: thisMonth(1),
    startingDayOfWeek: 1,
    displayPeriodUom: props.displayPeriodUom,
    items: [
        // {
        //     id: "e1",
        //     startDate: thisMonth(15, 18, 30),
        // },
        // {
        //     id: "e2",
        //     startDate: thisMonth(15),
        //     title: "Single-day item with a long title",
        // },
        // {
        //     id: "e3",
        //     startDate: thisMonth(7, 9, 25),
        //     endDate: thisMonth(10, 16, 30),
        //     title: "Multi-day item with a long title and times",
        // },
        // {
        //     id: "e4",
        //     startDate: thisMonth(20),
        //     title: "My Birthday!",
        //     classes: ["birthday"],
        //     url: "https://en.wikipedia.org/wiki/Birthday",
        // },
        // {
        //     id: "e5",
        //     startDate: thisMonth(5),
        //     endDate: thisMonth(12),
        //     title: "Multi-day item",
        //     classes: ["purple"],
        //     tooltip: "This spans multiple days",
        // },
        // {
        //     id: "foo",
        //     startDate: thisMonth(29),
        //     title: "Same day 1",
        // },
        // {
        //     id: "e6",
        //     startDate: thisMonth(29),
        //     title: "Same day 2",
        //     classes: ["orange"],
        // },
        // {
        //     id: "e7",
        //     startDate: thisMonth(29),
        //     title: "Same day 3",
        // },
        // {
        //     id: "e8",
        //     startDate: thisMonth(29),
        //     title: "Same day 4",
        //     classes: ["orange"],
        // },
        // {
        //     id: "e9",
        //     startDate: thisMonth(29),
        //     title: "Same day 5",
        // },
        // {
        //     id: "e10",
        //     startDate: thisMonth(29),
        //     title: "Same day 6",
        //     classes: ["orange"],
        // },
        // {
        //     id: "e11",
        //     startDate: thisMonth(29),
        //     title: "Same day 7",
        // },
    ],
});

const setShowDate = (d) => {
    data.value.showDate = d;
};

</script>
<style>
#app {
    /* color: #2c3e50; */
    height: 67vh;
    width: 90vw;
}
</style>
