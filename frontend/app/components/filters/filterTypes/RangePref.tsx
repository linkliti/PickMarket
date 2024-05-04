import FilterWrapper from "@/components/filters/filterTypes/FilterWrapper";
import { Input } from "@/components/ui/input";
import { Slider } from "@/components/ui/slider";
import { RangeFilter } from "@/proto/app/protos/types";

import { ReactElement } from "react";
import { Control, Controller } from "react-hook-form";

export default function RangePref({
  control,
  name,
  keyName,
  range,
}: {
  control: Control;
  name: string;
  keyName: string;
  range: RangeFilter;
}): ReactElement {
  return (
    <FilterWrapper
      name={name}
      keyName={keyName}
    >
      <Controller
        name={keyName}
        control={control}
        defaultValue=""
        render={({ field: { onChange, onBlur, value, disabled, name, ref } }) => (
          <div className="flex flex-row items-center gap-4 p-4">
            <Input
              className="mb-6 h-8 w-1/3 bg-white p-2"
              value={value}
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
                min={range.min}
                max={range.max}
                step={(range.max - range.min) / 100}
                value={[value]}
                onValueChange={(num: number[]): void => onChange(num[0])}
              />
              <p className="pt-2 text-left text-xs text-gray-500">
                {range.min}
                <span className="float-right">{range.max}</span>
              </p>
            </div>
          </div>
        )}
      />
    </FilterWrapper>
  );
}
