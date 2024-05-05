import { Category } from "@/proto/app/protos/categories";
import { Markets } from "@/proto/app/protos/types";
import { Marketplace } from "@/types/categoryTypes";
import { create } from "zustand";
import { devtools } from "zustand/middleware";

export type CategoryStore = {
  selectedMarket: Marketplace;
  selectedCategory: Category | null;
  setSelectedMarket: (market: Marketplace) => void;
  setSelectedCategory: (category: Category | null) => void;
};

export const marketplaces: Marketplace[] = [
  { label: "OZON", value: "https://ozon.ru", shortLabel: "ozon", id: Markets.OZON },
];

export const useCategoryStore = create(
  devtools<CategoryStore>((set) => ({
    selectedMarket: marketplaces[0],
    selectedCategory: null,
    setSelectedMarket: (market: Marketplace): void => set({ selectedMarket: market }),
    setSelectedCategory: (category: Category | null): void => set({ selectedCategory: category }),
  })),
);
