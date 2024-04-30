import { Category, CategoryStore, Marketplace } from "@/types/categoryTypes";
import { create } from "zustand";
import { devtools } from "zustand/middleware";

export const marketplaces: Marketplace[] = [
  { label: "OZON", value: "https://ozon.ru", shortLabel: "ozon" },
  { label: "Я.Маркет", value: "https://market.yandex.ru", shortLabel: "yand" },
];

export const useCategoryStore = create<CategoryStore>()(
  devtools((set) => ({
    selectedMarket: marketplaces[0],
    selectedCategory: null,
    setSelectedMarket: (market: Marketplace): void => set({ selectedMarket: market }),
    setSelectedCategory: (category: Category | null): void => set({ selectedCategory: category }),
  })),
);
