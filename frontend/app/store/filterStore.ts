import { ItemsRequestWithPrefs } from "@/proto/app/protos/reqHandlerTypes";
import { PrefForm } from "@/types/filterTypes";
import terminal from "virtual:terminal";
import { create } from "zustand";
import { devtools } from "zustand/middleware";

export const blacklistKeys: string[] = ["pm_isadult", "trucode", "sku"];

export type FilterStore = {
  activePrefs: ItemsRequestWithPrefs | null,
  setActivePrefs: (filters: ItemsRequestWithPrefs | null) => void,
  formPrefs: PrefForm | null,
  setFormPrefs: (filters: PrefForm | null) => void,
};

export const useFilterStore = create(
  devtools<FilterStore>((set) => ({
    activePrefs: null,
    setActivePrefs: (filters: ItemsRequestWithPrefs | null): void => {
      if (filters) {
        terminal.log(ItemsRequestWithPrefs.toJsonString(filters));
      }
      set({ activePrefs: filters })
    },
    formPrefs: null,
  setFormPrefs: (filters: PrefForm | null): void => set({ formPrefs: filters }),
  })),
);
