import { UserPref } from "@/proto/app/protos/reqHandlerTypes";
import { FiltersStore } from "@/types/filterTypes";
import { create } from "zustand";
import { devtools } from "zustand/middleware";
import { immer } from "zustand/middleware/immer";

export const blacklistKeys: string[] = ["pm_isAdult", "trucode", "sku"];

export const useFiltersStore = create(
  devtools(
    immer<FiltersStore>((set) => ({
      market: 0,
      pageUrl: "",
      numOfPages: 0,
      params: "",
      userQuery: "",
      prefs: {},
      setPageData: (market: number, pageUrl: string) => {
        set({ market: market, pageUrl: pageUrl });
      },
      setNumOfPages: (numOfPages: number) => {
        set({ numOfPages: numOfPages });
      },
      setParams: (params: string) => {
        set({ params: params });
      },
      setUserQuery: (userQuery: string) => {
        set(({ userQuery: userQuery }));
      },
      modifyPref: (key: string, value: UserPref) => {
        set((state) => {
          state.prefs[key] = value;
        });
      },
      resetStore: () => {
        set({
          market: 0,
          pageUrl: "",
          numOfPages: 0,
          params: "",
          userQuery: "",
          prefs: {},
        });
      },
    })),
  ),
);
