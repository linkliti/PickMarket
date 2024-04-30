import { Label } from "@/components/ui/label";
import { RadioGroupItem } from "@/components/ui/radio-group";
import { CategoryItemProps } from "@/types/categoryTypes";
import { ReactElement } from "react";

export default function CategoryItem({
  category,
  level,
  handleCategoryChange,
  selectedCategory,
}: CategoryItemProps): ReactElement {
  if (!category) {
    return <></>;
  }

  return (
    <div className="flex items-center space-x-2 py-0.5">
      <RadioGroupItem
        style={{ marginLeft: `${level * 1.6}rem` }}
        value={category.url}
        id={category.url}
        onClick={
          selectedCategory?.url !== category.url
            ? (): void => handleCategoryChange(category)
            : undefined
        }
        checked={category.url === selectedCategory?.url}
      />
      <Label
        htmlFor={category.url}
        className="cursor-pointer"
      >
        {category.title}
        {selectedCategory?.url === category.url ? " [выбрано]" : ""}
      </Label>
    </div>
  );
}
// <li
//   key={category.url}
//   style={{ paddingLeft: `${level}rem` }}
// >
//   <label>
//     <input
//       type="radio"
//       name="category"
//       value={category.url}
//       onClick={
//         selectedCategory?.url !== category.url
//           ? (): void => handleCategoryChange(category)
//           : undefined
//       }
//       defaultChecked={category.url === selectedCategory?.url}
//     />
//     {category.title}
//     {selectedCategory?.url === category.url ? " [выбрано]" : ""}
//   </label>
// </li>
