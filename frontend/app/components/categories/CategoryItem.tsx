import { CategoryItemProps } from "@/components/categories/types";
import { ReactElement } from "react";

export default function CategoryItem({
  category,
  level,
  handleCategoryChange,
  selectedCategory,
}: CategoryItemProps): ReactElement {
  return (
    <li key={category.url} style={{ paddingLeft: `${level}rem` }}>
      <label>
        <input
          type="radio"
          name="category"
          value={category.url}
          onClick={
            selectedCategory?.url !== category.url
              ? (): void => handleCategoryChange(category)
              : undefined
          }
          defaultChecked={category.url === selectedCategory?.url}
        />
        {category.title}
        {selectedCategory?.url === category.url ? ' [выбрано]' : ''}
      </label>
    </li>
  );
}
