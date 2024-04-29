import { useEffect, useState } from "react";
import { terminal } from "virtual:terminal";

type Category = {
  title: string;
  url: string;
  isParsed: boolean;
  children?: Category[];
}

export default function CategorySelect() {
  const [categories, setCategories] = useState<Category[]>([]);
  const [selectedCategory, setSelectedCategory] = useState<Category | null>(null);

  const updateCategory = (categories: Category[], url: string, children: Category[]): Category[] => {
    return categories.map(category => {
      if (category.url === url) {
        return { ...category, children: children.map(child => ({ ...child })), isParsed: true };
      }
      if (category.children) {
        return { ...category, children: updateCategory(category.children, url, children) };
      }
      return category;
    });
  };

  const fetchCategories = (url: string, parentCategory: Category | null = null) => {
    if (parentCategory && parentCategory.isParsed) {
      setSelectedCategory(parentCategory);
      terminal.log(parentCategory.title);
      return;
    }
    terminal.log('Загрузка категорий', url);
    fetch(url)
      .then(res => res.json())
      .then(data => {
        if (data !== null) {
          if (parentCategory) {
            setCategories(prevCategories => updateCategory(prevCategories, parentCategory.url, data));
          } else {
            setCategories(data.map((category: Category) => ({ ...category, isParsed: false })));
          }
        }
        setSelectedCategory(parentCategory);
        if (parentCategory) {
          terminal.log(parentCategory.title);
        }
      })
      .catch(terminal.error);
  };

  useEffect(() => {
    fetchCategories('/api/categories/ozon/root');
  }, []);

  const handleCategoryChange = (category: Category) => {
    fetchCategories(`/api/categories/ozon/sub?url=${encodeURIComponent(category.url)}`, category);
  };

  const renderCategories = (categories: Category[], level = 0) => {
    return categories.map(category => (
      <li
        key={category.url}
        style={{ paddingLeft: `${level * 1}rem` }}
      >
        <label>
          <input
            type="radio"
            name="category"
            value={category.url}
            onChange={() => handleCategoryChange(category)}
            defaultChecked={category.url === '/'}
          />
          {category.title}
          {selectedCategory && selectedCategory.url === category.url && ' [выбрано]'}
        </label>
        {category.children && <ul>{renderCategories(category.children, level + 1)}</ul>}
      </li>
    ));
  };

  return (
    <form>
      <ul>
        {renderCategories(categories)}
      </ul>
      {selectedCategory && <p>Выбранная категория: {selectedCategory.title}</p>}
    </form>
  );
}