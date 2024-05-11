import { ReactElement } from "react";
import { Link } from "react-router-dom";

type navItemsType = {
  name: string;
  path: string;
};

const navItems: navItemsType[] = [
  { name: "Категории", path: "/categories" },
  { name: "Избранное", path: "/favorites" },
];

export default function Header(): ReactElement {
  return (
    <header className="inline-flex w-full flex-wrap items-center rounded-2xl bg-sky-200 p-4">
      <div className="grow">
        <Link
          to="/"
          className="text-2xl font-extrabold"
        >
          PickMarket
        </Link>
      </div>
      <nav className="inline-flex flex-wrap gap-x-5 gap-y-2 ">
        {navItems.map(
          (item: navItemsType): ReactElement => (
            <Link
              key={item.name}
              to={item.path}
              className="font-bold"
            >
              {item.name}
            </Link>
          ),
        )}
      </nav>
    </header>
  );
}
