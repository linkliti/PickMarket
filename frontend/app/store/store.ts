import { create } from "zustand";

type CouterStore = {
  count: number;
}

export const useStore = create<CouterStore>(() => ({
  count: 0,
}))