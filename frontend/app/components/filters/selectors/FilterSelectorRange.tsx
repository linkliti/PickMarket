import { Input } from "@/components/ui/input";
import { Slider } from "@/components/ui/slider";
import { Filter } from "@/proto/app/protos/items";
import { useFiltersStore } from "@/store/filtersStore";
import { FilterPrefValue, FiltersStore } from "@/types/filterTypes";

import { ReactElement } from "react";
import terminal from "virtual:terminal";

export default function FilterSelectorRange({ filter }: { filter: Filter }): ReactElement {
  const [prefs, modifyPref] = useFiltersStore((state: FiltersStore) => [
    state.prefs,
    state.modifyPref,
  ]);

  function updateRangeSelector(num: number): void {
    const v: FilterPrefValue = prefs[filter.key].value;
    if (v.oneofKind === "numVal") {
      modifyPref(filter.key, {
        priority: prefs[filter.key].priority,
        value: {
          oneofKind: "numVal",
          numVal: num,
        },
      });
      terminal.log("Updated filter", filter.key, num);
    }
  }

  function getNumValue(): number {
    const v: FilterPrefValue = prefs[filter.key].value;
    if (v.oneofKind === "numVal") {
      return v.numVal;
    }
    return 0;
  }

  if (filter.data.oneofKind !== "rangeFilter") return <>WrongType!</>;
  const min = filter.data.rangeFilter.min;
  const max = filter.data.rangeFilter.max;

  return (
    <div className="flex flex-row items-center gap-4 p-4">
      <Input
        value={getNumValue()}
        className="mb-6 h-8 w-1/3 bg-white p-2"
        onChange={(event: React.ChangeEvent<HTMLInputElement>): void => {
          const num: number = parseInt(event.target.value, 10);
          if (!isNaN(num)) updateRangeSelector(num);
        }}
      />
      <div className="flex grow flex-col">
        <Slider
          className=""
          value={[getNumValue()]}
          onValueChange={(num: number[]): void => updateRangeSelector(num[0])}
          max={max}
          min={min}
        />
        <p className="pt-2 text-left text-xs text-gray-500">
          {min}
          <span className="float-right">{max}</span>
        </p>
      </div>
    </div>
  );
}
