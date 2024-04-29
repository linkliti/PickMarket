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
