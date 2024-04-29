import { Link } from "react-router-dom";

export default function NotFound() {
  return (
    <section className="mt-4 flex w-full grow flex-col items-center justify-center rounded-2xl bg-white  p-4 max-md:max-w-full">
      <div className="mx-auto max-w-xl text-center">
        <h1 className="text-4xl font-extrabold">404</h1>

        <p className="mt-4 text-lg">Страница не найдена :(</p>

        <div className="mt-8 flex flex-wrap justify-center gap-4">
          <Link to="/">
            <a className="text-md block w-full rounded-lg bg-sky-200 px-5 py-3 font-medium text-black sm:w-auto">
              Перейти на главную
            </a>
          </Link>
        </div>
      </div>
    </section>
  );
}
