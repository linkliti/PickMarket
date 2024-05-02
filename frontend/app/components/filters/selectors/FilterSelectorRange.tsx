import { Input } from "@/components/ui/input";
import { Slider } from "@/components/ui/slider";
import { Filter } from "@/proto/app/protos/items";

import { ReactElement } from "react";

export default function FilterSelectorRange({ filter }: { filter: Filter }): ReactElement {
  if (filter.data.oneofKind !== "rangeFilter") return <>WrongType!</>;
  const min = filter.data.rangeFilter.min;
  const max = filter.data.rangeFilter.max;

  return (
    <div className="flex flex-row items-center gap-4 p-4">
      <Input
        className="mb-6 h-8 w-1/3 bg-white p-2"
        defaultValue={max}
      />
      <div className="flex grow flex-col">
        <Slider
          className=""
          defaultValue={[max]}
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
