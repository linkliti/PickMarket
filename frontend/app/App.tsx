import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Header from "./components/base/Header";

interface CategoryItemProps {
  icon: string;
  label: string;
}

const CategoryItem: React.FC<CategoryItemProps> = ({ icon, label }) => (
  <div className="justify-center py-2 mt-2.5 bg-white rounded-md border border-black border-solid">
    {icon} {label}
  </div>
);

interface SubCategoryItemProps {
  label: string;
  isActive?: boolean;
}

const SubCategoryItem: React.FC<SubCategoryItemProps> = ({
  label,
  isActive = false,
}) => (
  <div
    className={`flex gap-5 px-1.5 py-2 mt-3.5 rounded-md border border-black border-solid ${
      isActive ? "bg-neutral-200" : ""
    }`}
  >
    <div className="flex-auto">🏠 {label}</div>
    {isActive && <div className="my-auto">&gt;</div>}
  </div>
);

const categories: CategoryItemProps[] = [
  { icon: "📦", label: "Категория 2" },
  { icon: "🪑", label: "Категория 3" },
];

const subCategories: SubCategoryItemProps[] = [
  { label: "Категория 1", isActive: true },
  { label: "Категория 2" },
  { label: "Категория 3" },
];

const router = createBrowserRouter([
  {
    path: "/",
    element: <Base />,
  },
]);

function Base() {
  return (
    <>
      <section className="flex flex-col px-4 pt-2.5 pb-5 mt-4 font-bold bg-white rounded-2xl max-md:max-w-full">
        <div className="flex gap-5 w-full whitespace-nowrap max-md:flex-wrap max-md:max-w-full">
          <h1 className="flex-auto my-auto text-2xl text-black">
            Маркетплейс:
          </h1>
          <div className="flex gap-4 text-lg text-center text-black">
            <div className="justify-center px-9 py-3.5 bg-sky-200 rounded-2xl max-md:px-5">
              OZON
            </div>
            <div className="justify-center px-5 py-3 bg-white rounded-2xl border border-black border-solid">
              Я.Маркет
            </div>
          </div>
          <form className="flex gap-5 justify-between">
            <label
              htmlFor="search"
              className="my-auto text-lg text-center text-black"
            >
              Поиск:
            </label>
            <input
              type="text"
              id="search"
              placeholder="Поиск"
              className="justify-center py-4 text-base bg-white rounded-2xl border border-black border-solid text-zinc-500"
            />
          </form>
        </div>
        <div className="flex gap-5 mt-5 text-lg text-black max-md:flex-wrap max-md:max-w-full">
          <div className="flex-auto my-auto max-md:max-w-full">
            Выбранная категория: "Категория - ДочерКатегория"
          </div>
          <button className="justify-center px-12 py-3.5 text-center bg-sky-200 rounded-2xl max-md:px-5">
            Показать товары
          </button>
        </div>
      </section>
      <div className="px-px mt-4 max-md:max-w-full">
        <div className="flex gap-5 max-md:flex-col max-md:gap-0">
          <aside className="flex flex-col w-3/12 max-md:ml-0 max-md:w-full">
            <div className="flex flex-col grow px-4 pt-3 pb-20 mx-auto w-full text-base text-black bg-white rounded-2xl max-md:mt-4">
              <h2 className="text-2xl font-bold">Категории</h2>
              <SubCategoryItem label="Категория 1" isActive />
              {categories.map((category) => (
                <CategoryItem
                  key={category.label}
                  icon={category.icon}
                  label={category.label}
                />
              ))}
            </div>
          </aside>
          <aside className="flex flex-col ml-5 w-3/12 max-md:ml-0 max-md:w-full">
            <div className="flex flex-col grow px-4 pt-3 pb-20 mx-auto w-full text-base text-black bg-white rounded-2xl max-md:mt-4">
              <h2 className="text-2xl font-bold">Подкатегории</h2>
              {subCategories.map((subCategory) => (
                <SubCategoryItem
                  key={subCategory.label}
                  label={subCategory.label}
                  isActive={subCategory.isActive}
                />
              ))}
            </div>
          </aside>
          <main className="flex flex-col ml-5 w-[51%] max-md:ml-0 max-md:w-full">
            <div className="flex flex-col grow items-start pt-3 pr-20 pb-20 pl-4 w-full text-black bg-white rounded-2xl max-md:pr-5 max-md:mt-4 max-md:max-w-full">
              <h2 className="text-2xl font-bold">Дочерние категории</h2>
              <div className="mt-6 text-base leading-8">
                <ul>
                  <li>
                    <span className="leading-7">⬇️ </span>Уровень 3{" "}
                  </li>
                  <ul>
                    <ul>
                      <li>Уровень 4</li>
                      <li>Уровень 4</li>
                    </ul>
                  </ul>
                  <li>
                    <span className="leading-7">⬆️ </span>Уровень 3{" "}
                  </li>
                  <li>
                    <span className="leading-7">⬇️ </span>Уровень 3{" "}
                  </li>
                  <ul>
                    <ul>
                      <li>
                        <span className="leading-7">⬇️ </span>Уровень 4{" "}
                      </li>
                      <ul>
                        <ul>
                          <li>Уровень 5</li>
                        </ul>
                      </ul>
                      <li>Уровень 4</li>
                    </ul>
                  </ul>
                </ul>
                <br />
              </div>
            </div>
          </main>
        </div>
      </div>
    </>
  );
}

function App() {
  return (
    <div className="min-h-screen flex flex-col items-center px-16 pt-10 pb-20 bg-sky-100 max-md:px-5">
      <div className="flex flex-col w-full max-w-[1065px] max-md:max-w-full">
        <Header />
        <RouterProvider router={router} />
      </div>
    </div>
  );
}

export default App;
