<template>
  <div class="font-sans text-gray-900">
    <!-- Header -->
    <div class="grid sticky top-0 bg-white z-20 border-b border-gray-200"
:style="{ gridTemplateColumns: `150px repeat(${days.length}, 1fr)` }">
      <div class="h-11 border-r border-gray-200"></div>
      <div v-for="d in days" :key="d" class="grid grid-flow-col auto-cols-fr place-items-center h-11 font-semibold">
        {{ d }}
      </div>
    </div>

    <!-- Body -->
    <div class="grid grid-cols-[100px_1fr]" >
      <!-- Time column -->
      <div class="border-r border-gray-200">
        <div
          v-for="t in gridTimes"
          :key="t.key"
          class="flex items-start justify-end pr-2 text-xs text-gray-500 box-border"
          :style="{ height: rowHeight + 'px' }"
        >
          <span v-if="t.showLabel">{{ t.label }}</span>
        </div>
      </div>

      <!-- Days grid -->
      <div class="relative grid" :style="{ gridTemplateColumns: `repeat(${days.length}, 1fr)` }">
        <!-- background rows -->
        <template v-for="t in gridTimes" :key="t.key + '-bg'">
          <div
            v-for="i in days.length"
            :key="t.key + '-bg-' + i"
            class="border-b border-dashed border-gray-200"
            :style="{ height: rowHeight + 'px' }"
          />
        </template>

        <!-- events -->
        <div
          v-for="ev in laidOutEvents"
          :key="ev.id"
          class="absolute box-border rounded-lg overflow-hidden backdrop-blur-sm p-2.5 pl-3"
          :style="eventStyle(ev)"
          @mouseenter="hoverId = ev.id"
          @mouseleave="hoverId = null"
        >
          <div class="font-semibold text-xs leading-tight">
            {{ ev.subject }}
          </div>
          <div class="text-[11px] text-gray-700 mt-0.5">
            <span v-if="ev.room">Aula {{ ev.room }}</span>
            <span v-if="ev.teacher">• {{ ev.teacher }}</span>
          </div>

          <!-- tooltip -->
          <div
            v-if="hoverId === ev.id"
            class="absolute bottom-1 left-1 bg-gray-900 text-white rounded-md p-2 text-xs shadow-lg w-max max-w-xs"
          >
            <div class="font-bold mb-0.5">{{ ev.subject }}</div>
            <div>{{ prettyTime(ev.startDate) }} – {{ prettyTime(ev.endDate) }}</div>
            <div v-if="ev.room">Aula: {{ ev.room }}</div>
            <div v-if="ev.teacher">Docente: {{ ev.teacher }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Legend -->
    <!-- <div v-if="legend && legend.length" class="mt-3 flex flex-wrap gap-2.5">
      <div
        v-for="item in legend"
        :key="item.label"
        class="inline-flex items-center gap-1.5 text-xs text-gray-700"
      >
        <span class="w-2.5 h-2.5 rounded-full inline-block" :style="{ background: item.color }"></span>
        {{ item.label }}
      </div>
    </div> -->
  </div>
</template>

<script setup lang="ts">

import { computed, ref, toRaw } from 'vue';
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import timezone from "dayjs/plugin/timezone";

dayjs.extend(utc);
dayjs.extend(timezone);

const TIMEZONE = "America/Lima";

type EventItem = {
  id: string | number;
  DayIndex: number; // 0..days.length-1
  subject: string;
  startDate: string; // "HH:MM"
  endDate: string;   // "HH:MM"
  room?: string | number;
  teacher?: string;
  color?: string;
};

const props = withDefaults(defineProps<{
  days: string[];
  startHour: string;
  endHour: string;
  slotMinutes?: number;
  rowHeight?: number;
  events: EventItem[];
  legend?: { label: string; color: string }[];
}>(), {
  slotMinutes: 30,
  rowHeight: 36,
  legend: () => []
});

const hoverId = ref<string | number | null>(null);

function toMinutesFromString(iso: string): number {
  const [h, m] = iso.split(':').map(Number);
  return h * 60 + m;
}

function toMinutes(iso: string): number {
  const hhmm = iso.substring(11, 16); // "20:00"
  const [h, m] = hhmm.split(':').map(Number);
  return h * 60 + m;
}
const dayStart = computed(() => toMinutesFromString(props.startHour));
const dayEnd = computed(() => toMinutesFromString(props.endHour));
const totalMinutes = computed(() => dayEnd.value - dayStart.value);
const rows = computed(() => Math.ceil(totalMinutes.value / props.slotMinutes));
const rowHeight = computed(() => props.rowHeight);

const gridTimes = computed(() => {
  // Muestra etiqueta cada hora “redonda”
  const out: { key: string; label: string; showLabel: boolean }[] = [];
  for (let i = 0; i < rows.value; i++) {
    const minutesFromStart = i * props.slotMinutes + dayStart.value;
    const h = Math.floor(minutesFromStart / 60);
    const m = minutesFromStart % 60;
    const label = `${String(h).padStart(2,'0')}:${String(m).padStart(2,'0')}`;
    const showLabel = m === 0; // etiqueta cada 60 min
    out.push({ key: `t-${i}`, label, showLabel });
  }
  return out;
});

// Layout de eventos: calcula top/height y columna; maneja choques con "lanes"
type LaidOut = EventItem & {
  top: number;        // px
  height: number;     // px
  col: number;        // 0..days-1
  lane: number;       // carril por choque
  lanesInColSlot: number; // total de carriles para ese intervalo y día
};

const laidOutEvents = computed<LaidOut[]>(() => {
  // Agrupar por día
  const byDay: Record<number, EventItem[]> = {};
  console.log('PROPS.EVENT', props.events)
  const rawEvents = toRaw(props.events)
  rawEvents.forEach(ev => {
    if (ev.DayIndex < 0 || ev.DayIndex >= props.days.length) return;
    (byDay[ev.DayIndex] ||= []).push(ev);
  });

  const result: LaidOut[] = [];

  console.log('BYDAYS', byDay)
  for (let d = 0; d < props.days.length; d++) {
    // console.log('d', d)
    // const evs = (byDay[d] || []).slice().sort((a, b) => toMinutes(a.endDate) - toMinutes(b.startDate));
    const evs = (byDay[d] || [])
    .map(ev => ({
      ...ev,
      startDate: dayjs.utc(ev.startDate).tz(TIMEZONE).format("YYYY-MM-DDTHH:mm:ss"),
      endDate: dayjs.utc(ev.endDate).tz(TIMEZONE).format("YYYY-MM-DDTHH:mm:ss"),
    }))
    .slice()
    .sort((a, b) => toMinutes(a.endDate) - toMinutes(b.startDate));

    console.log('EVS -- UPDATE', evs)

    // Algoritmo de lanes por choques (line sweep simple)
    type Active = { end: number; lane: number };
    const active: Active[] = [];
    let maxLaneUsed = 0;

    const pushActive = (end: number, lane: number) => active.push({ end, lane });
    const releaseFinished = (currentStart: number) => {
      for (let i = active.length - 1; i >= 0; i--) {
        if (active[i].end <= currentStart) active.splice(i, 1);
      }
    };
    const nextFreeLane = (): number => {
      const used = new Set(active.map(a => a.lane));
      let lane = 0;
      while (used.has(lane)) lane++;
      return lane;
    };


    const intervals: { start: number; end: number; lanes: number }[] = [];

    evs.forEach(ev => {
      const s = toMinutes(ev.startDate);
      const e = toMinutes(ev.endDate);
      releaseFinished(s);
      const lane = nextFreeLane();
      pushActive(e, lane);
      maxLaneUsed = Math.max(maxLaneUsed, lane);

      // hallar cuántos lanes simultáneos hay en este punto
      const lanesNow = Math.max(...active.map(a => a.lane), 0) + 1;
      intervals.push({ start: s, end: e, lanes: lanesNow });

      const topMin = s - dayStart.value;
      const durMin = e - s;
      const topPx = (topMin / props.slotMinutes) * rowHeight.value;
      const heightPx = (durMin / props.slotMinutes) * rowHeight.value;

      result.push({
        ...ev,
        top: topPx,
        height: Math.max(heightPx, 0), // evita negativos
        col: d,
        lane,
        lanesInColSlot: lanesNow
      });
    });

    // Normaliza lanesInColSlot por cluster
    result
      .filter(r => r.col === d)
      .forEach(r => {
        const maxLanes = intervals
          .filter(i => !(toMinutes(r.endDate) <= i.start || toMinutes(r.startDate) >= i.end))
          .reduce((mx, i) => Math.max(mx, i.lanes), 1);
        r.lanesInColSlot = Math.max(1, maxLanes);
      });
  }

  return result;
});

function eventStyle(ev: LaidOut) {
  const colWidthPercent = 100 / props.days.length;
  const leftBase = ev.col * colWidthPercent;
  const laneWidth = colWidthPercent / ev.lanesInColSlot;
  const left = leftBase + ev.lane * laneWidth;

  return {
    top: ev.top + 'px',
    height: ev.height + 'px',
    left: left + '%',
    width: laneWidth + '%',
    background: ev.color || 'rgba(59,130,246,0.12)',  // por defecto azul suave
    borderLeft: `4px solid ${ev.color || '#3b82f6'}`
  };
}

function prettyTime(isoDate: string) {
  if (!isoDate) return "";
  const date = new Date(isoDate);
  return date.toLocaleTimeString("es-PE", {
    hour: "2-digit",
    minute: "2-digit",
    hour12: false, // <-- fuerza formato 24h
  });
}

</script>