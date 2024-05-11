import { Category } from "@/proto/app/protos/categories";

// Parents array of selectedCategory for nesting
export default function findParents(parentUrl: string, categories: Category[]): Category[] {
  const parentCategories: Category[] = [];
  for (;;) {
    const parentCat: Category = categories.filter(
      (category: Category): boolean => parentUrl === category.url,
    )[0];
    parentCategories.unshift(parentCat);
    if (parentCat.parentUrl) {
      parentUrl = parentCat.parentUrl;
    } else {
      break;
    }
  }
  return parentCategories;
}