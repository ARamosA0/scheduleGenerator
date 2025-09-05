<template>
  <div class="schedule">
    <!-- Header: days -->
    <div class="header">
      <div class="corner"></div>
      <div v-for="d in days" :key="d" class="day">{{ d }}</div>
    </div>

    <!-- Body: time grid + events -->
    <div class="body">
      <!-- time column -->
      <div class="times">
        <div
          v-for="t in gridTimes"
          :key="t.key"
          class="time-cell"
          :style="{ height: rowHeight + 'px' }"
        >
          <span v-if="t.showLabel">{{ t.label }}</span>
        </div>
      </div>

      <!-- days grid -->
      <div class="days-grid" :style="{ gridTemplateColumns: `repeat(${days.length}, 1fr)` }">
        <!-- background rows -->
        <template v-for="t in gridTimes" :key="t.key + '-bg'">
          <div
            class="bg-row"
            :style="{ height: rowHeight + 'px' }"
            v-for="i in days.length"
            :key="t.key + '-bg-' + i"
          />
        </template>

        <!-- events -->
        <div
          v-for="ev in laidOutEvents"
          :key="ev.id"
          class="event"
          :style="eventStyle(ev)"
          @mouseenter="hoverId = ev.id"
          @mouseleave="hoverId = null"
        >
          <div class="event-title">{{ ev.title }}</div>
          <div class="event-meta">
            <span v-if="ev.room">Aula {{ ev.room }}</span>
            <span v-if="ev.teacher">• {{ ev.teacher }}</span>
          </div>

          <!-- tooltip -->
          <div class="tooltip" v-if="hoverId === ev.id">
            <div class="tooltip-title">{{ ev.title }}</div>
            <div>{{ prettyTime(ev.start) }} – {{ prettyTime(ev.end) }}</div>
            <div v-if="ev.room">Aula: {{ ev.room }}</div>
            <div v-if="ev.teacher">Docente: {{ ev.teacher }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Legend (optional) -->
    <div v-if="legend && legend.length" class="legend">
      <div v-for="item in legend" :key="item.label" class="legend-item">
        <span class="dot" :style="{ background: item.color }"></span>{{ item.label }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
/**
 * Props
 * - days: nombres de días (e.g., ["Lun","Mar","Mié","Jue","Vie"])
 * - startHour / endHour: límites del día en formato "HH:MM"
 * - slotMinutes: tamaño de cada fila del grid en minutos (e.g., 30)
 * - rowHeight: alto visual de cada fila (px)
 * - events: arreglo con clases:
 *   {
 *     id: string|number,
 *     dayIndex: number (0 = primer día),
 *     title: string,
 *     start: "HH:MM",
 *     end: "HH:MM",
 *     room?: string|number,
 *     teacher?: string,
 *     color?: string   // opcional, color de fondo del evento
 *   }
 * - legend: [{ label: string, color: string }]
 */
import { computed, ref } from 'vue';

type EventItem = {
  id: string | number;
  dayIndex: number; // 0..days.length-1
  title: string;
  start: string; // "HH:MM"
  end: string;   // "HH:MM"
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

function toMinutes(hhmm: string): number {
  const [h, m] = hhmm.split(':').map(Number);
  return h * 60 + m;
}
const dayStart = computed(() => toMinutes(props.startHour));
const dayEnd = computed(() => toMinutes(props.endHour));
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
  props.events.forEach(ev => {
    if (ev.dayIndex < 0 || ev.dayIndex >= props.days.length) return;
    (byDay[ev.dayIndex] ||= []).push(ev);
  });

  const result: LaidOut[] = [];

  for (let d = 0; d < props.days.length; d++) {
    const evs = (byDay[d] || []).slice().sort((a, b) => toMinutes(a.start) - toMinutes(b.start));

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

    // Para conocer cuántos carriles ocupa cada “cluster”
    // mapeamos intervalos a su máximo de lanes
    const intervals: { start: number; end: number; lanes: number }[] = [];

    evs.forEach(ev => {
      const s = toMinutes(ev.start);
      const e = toMinutes(ev.end);
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
          .filter(i => !(toMinutes(r.end) <= i.start || toMinutes(r.start) >= i.end))
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

function prettyTime(hhmm: string) {
  return hhmm;
}
</script>

<style scoped>
.schedule {
  font-family: system-ui, -apple-system, Segoe UI, Roboto, Ubuntu, Cantarell, 'Helvetica Neue', Arial;
  color: #111827;
}

/* Header */
.header {
  display: grid;
  grid-template-columns: 90px 1fr;
  gap: 0;
  position: sticky;
  top: 0;
  background: #fff;
  z-index: 2;
  border-bottom: 1px solid #e5e7eb;
}
.header .corner {
  height: 44px;
  border-right: 1px solid #e5e7eb;
}
.header .day {
  display: grid;
  grid-auto-flow: column;
  grid-auto-columns: 1fr;
  align-items: center;
  justify-items: center;
  height: 44px;
  font-weight: 600;
}

/* Body */
.body {
  display: grid;
  grid-template-columns: 90px 1fr;
}
.times {
  border-right: 1px solid #e5e7eb;
}
.time-cell {
  display: flex;
  align-items: flex-start;
  justify-content: flex-end;
  padding-right: 8px;
  font-size: 12px;
  color: #6b7280;
  box-sizing: border-box;
}
.days-grid {
  position: relative;
  display: grid;
}
.bg-row {
  border-bottom: 1px dashed #e5e7eb;
}

/* Events */
.event {
  position: absolute;
  box-sizing: border-box;
  padding: 6px 8px 6px 10px;
  border-radius: 8px;
  overflow: hidden;
  backdrop-filter: blur(1px);
}
.event-title {
  font-weight: 600;
  font-size: 12px;
  line-height: 1.2;
}
.event-meta {
  font-size: 11px;
  color: #374151;
  margin-top: 2px;
}

/* Tooltip */
.tooltip {
  position: absolute;
  left: 8px;
  right: 8px;
  bottom: 8px;
  background: #111827;
  color: #fff;
  border-radius: 6px;
  padding: 6px 8px;
  font-size: 12px;
  box-shadow: 0 8px 20px rgba(0,0,0,0.18);
}
.tooltip-title {
  font-weight: 700;
  margin-bottom: 2px;
}

/* Legend */
.legend {
  margin-top: 12px;
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}
.legend-item {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: #374151;
}
.legend-item .dot {
  width: 10px;
  height: 10px;
  border-radius: 999px;
  display: inline-block;
}
</style>