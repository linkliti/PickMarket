import MultipleSelector, { Option } from "@/components/ui/multiple-selector";
import { Filter } from "@/proto/app/protos/items";
import { useFiltersStore } from "@/store/filtersStore";
import { FilterPrefValue, FiltersStore } from "@/types/filterTypes";

import { ReactElement } from "react";
import terminal from "virtual:terminal";

export default function FilterSelectorList({ filter }: { filter: Filter }): ReactElement {
  const [prefs, modifyPref] = useFiltersStore((state: FiltersStore) => [
    state.prefs,
    state.modifyPref,
  ]);

  if (filter.data.oneofKind !== "selectionFilter") return <>WrongType!</>;
  const options: Option[] = [];

  for (const filt of filter.data.selectionFilter.items) {
    options.push({ value: filt.text, label: filt.text });
  }

  function updateListSelector(options: Option[]): void {
    const v: FilterPrefValue = prefs[filter.key].value;
    if (v.oneofKind === "listVal") {
      const newVals = options.map((i: Option): string => i.value);
      modifyPref(filter.key, {
        priority: prefs[filter.key].priority,
        value: {
          oneofKind: "listVal",
          listVal: { values: newVals },
        },
      });
      terminal.log("Updated filter", filter.key, newVals);
    }
  }

  function getListValues(): Option[] {
    const v: FilterPrefValue = prefs[filter.key].value;
    if (v.oneofKind === "listVal") {
      return v.listVal.values.map((i: string): Option => ({ value: i, label: i }));
    }
    return [];
  }

  return (
    <>
      <MultipleSelector
        className="mx-6 mt-2 bg-white"
        defaultOptions={options}
        value={getListValues()}
        onChange={updateListSelector}
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
