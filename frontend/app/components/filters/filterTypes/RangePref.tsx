import FilterWrapper from "@/components/filters/filterTypes/FilterWrapper";
import { Input } from "@/components/ui/input";
import { Slider } from "@/components/ui/slider";
import { RangeFilter } from "@/proto/app/protos/types";
import { PrefForm } from "@/types/filterTypes";

import { ReactElement } from "react";
import { Control, Controller } from "react-hook-form";

export default function RangePref({
  control,
  filterTitle,
  keyName,
  range,
}: {
  control: Control<PrefForm, unknown>;
  filterTitle: string;
  keyName: string;
  range: RangeFilter;
}): ReactElement {
  return (
    <Controller
      name={`prefs.${keyName}`}
      control={control}
      defaultValue={range.max}
      render={({ field: { onChange, onBlur, value, disabled, name, ref } }) => {
        // Transform value to always be a number or undefined
        const transformedValue: number = typeof value === "number" ? value : 0;

        return (
          <FilterWrapper
            name={filterTitle}
            keyName={keyName}
          >
            <div className="flex flex-row items-center gap-4 p-4">
              <Input
                className="mb-6 h-8 w-1/3 bg-white p-2"
                value={transformedValue}
                onChange={(event: React.ChangeEvent<HTMLInputElement>): void => {
                  const num: number = parseInt(event.target.value, 10);
                  if (isNaN(num) || num === 0) {
                    onChange(range.min);
                  } else {
                    onChange(num);
                  }
                }}
                ref={ref}
                onBlur={onBlur}
                disabled={disabled}
                name={name}
              />
              <div className="flex grow flex-col">
                <Slider
                  className=""
                  min={range.min}
                  max={range.max}
                  step={(range.max - range.min) / 100}
                  value={[transformedValue]}
                  onValueChange={(num: number[]): void => onChange(num[0])}
                />
                <p className="pt-2 text-left text-xs text-gray-500">
                  {range.min}
                  <span className="float-right">{range.max}</span>
                </p>
              </div>
            </div>
          </FilterWrapper>
        );
      }}
    />
  );
}
