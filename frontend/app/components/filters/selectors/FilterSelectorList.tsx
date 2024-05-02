import MultipleSelector, { Option } from "@/components/ui/multiple-selector";
import { Filter } from "@/proto/app/protos/items";

import { ReactElement } from "react";

export default function FilterSelectorList({ filter }: { filter: Filter }): ReactElement {
  if (filter.data.oneofKind !== "selectionFilter") return <>WrongType!</>;

  const options: Option[] = [];

  for (const filt of filter.data.selectionFilter.items) {
    options.push({ value: filt.text, label: filt.text });
  }

  return (
    <>
      <MultipleSelector
        className="mx-6 mt-2 bg-white"
        defaultOptions={options}
        placeholder="Выберите предпочтения..."
        emptyIndicator={
          <p className="text-center text-lg leading-10 text-gray-500 dark:text-gray-500">
            Не найдено
          </p>
        }
      />
    </>
  );
}
