import { Category } from "@/proto/app/protos/categories";
import { Markets } from "@/proto/app/protos/types";

export type CategoryItemProps = {
  category: Category;
  level: number;
  handleCategoryChange: (category: Category) => void;
  selectedCategory?: Category;
};

export type Marketplace = {
  label: string;
  value: string;
  shortLabel: string;
  id: Markets
};
