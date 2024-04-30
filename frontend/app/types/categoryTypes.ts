export type Category = {
  title: string;
  url: string;
  parentUrl?: string;
  isParsed: boolean;
};

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
};

export type CategoryStore = {
  selectedMarket: Marketplace;
  selectedCategory: Category | null;
  setSelectedMarket: (market: Marketplace) => void;
  setSelectedCategory: (category: Category | null) => void;
};

