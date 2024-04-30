import { Marketplace } from "@/types/categoryTypes";
import { create } from "zustand";
import { devtools } from "zustand/middleware";

export const marketplaces: Marketplace[] = [
  { label: "OZON", value: "https://ozon.ru" },
  { label: "Я.Маркет", value: "https://market.yandex.ru" },
];

type CategoryStore = {
  selectedMarket: Marketplace;
  setSelectedMarket: (market: Marketplace) => void;
  selectedCategory: string | null;
  setSelectedCategory: (category: string) => void;
};

export const categoryStore = create<CategoryStore>()(
  devtools((set) => ({
    selectedMarket: marketplaces[0],
    setSelectedMarket: (market: Marketplace): void => set({ selectedMarket: market }),
    selectedCategory: null,
    setSelectedCategory: (category: string): void => set({ selectedCategory: category }),
  })),
);
