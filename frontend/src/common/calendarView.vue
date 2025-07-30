<template>
    <div id="app">
        <CalendarView
            :show-date="data.showDate"
            :items="data.items"
            :starting-day-of-week="data.startingDayOfWeek"
            :display-period-uom="props.displayPeriodUom"
            class="theme-default"
        >
            <template #header="{ headerProps }">
                <CalendarViewHeader
                    slot="header"
                    :header-props="headerProps"
                    @input="data.showDate"
                />
            </template>
        </CalendarView>
    </div>
</template>
<script setup lang="js">
import { ref } from "vue";
import { CalendarView, CalendarViewHeader } from "vue-simple-calendar";

import "../../node_modules/vue-simple-calendar/dist/vue-simple-calendar.css";
import "../../node_modules/vue-simple-calendar/dist/css/default.css";

const props = defineProps({
    displayPeriodUom: {
        type: Object,
        deault: null,
    },
});

const thisMonth = (d, h, m) => {
    const t = new Date();
    return new Date(t.getFullYear(), t.getMonth(), d, h ?? 0, m ?? 0);
};

const data = ref({
    showDate: thisMonth(1),
    startingDayOfWeek: 0,
    displayPeriodUom: props.displayPeriodUom,
    items: [
        {
            id: "e1",
            startDate: thisMonth(15, 18, 30),
        },
        {
            id: "e2",
            startDate: thisMonth(15),
            title: "Single-day item with a long title",
        },
        {
            id: "e3",
            startDate: thisMonth(7, 9, 25),
            endDate: thisMonth(10, 16, 30),
            title: "Multi-day item with a long title and times",
        },
        {
            id: "e4",
            startDate: thisMonth(20),
            title: "My Birthday!",
            classes: ["birthday"],
            url: "https://en.wikipedia.org/wiki/Birthday",
        },
        {
            id: "e5",
            startDate: thisMonth(5),
            endDate: thisMonth(12),
            title: "Multi-day item",
            classes: ["purple"],
            tooltip: "This spans multiple days",
        },
        {
            id: "foo",
            startDate: thisMonth(29),
            title: "Same day 1",
        },
        {
            id: "e6",
            startDate: thisMonth(29),
            title: "Same day 2",
            classes: ["orange"],
        },
        {
            id: "e7",
            startDate: thisMonth(29),
            title: "Same day 3",
        },
        {
            id: "e8",
            startDate: thisMonth(29),
            title: "Same day 4",
            classes: ["orange"],
        },
        {
            id: "e9",
            startDate: thisMonth(29),
            title: "Same day 5",
        },
        {
            id: "e10",
            startDate: thisMonth(29),
            title: "Same day 6",
            classes: ["orange"],
        },
        {
            id: "e11",
            startDate: thisMonth(29),
            title: "Same day 7",
        },
    ],
});

const setShowDate = (d) => {
    showDate.value = d;
};
</script>
<style>
#app {
    /* color: #2c3e50; */
    height: 67vh;
    width: 90vw;
}
</style>
