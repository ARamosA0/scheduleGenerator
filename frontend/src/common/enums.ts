export const RoomTypes = {
  Classroom: 1,
  Laboratory: 2,
  Auditorium: 3,
  Office: 4,
} as const;

export type RoomTypes = (typeof RoomTypes)[keyof typeof RoomTypes];
