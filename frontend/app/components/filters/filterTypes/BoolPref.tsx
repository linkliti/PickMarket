import PrioritySelector from "@/components/filters/filterTypes/PrioritySelector";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import { cn } from "@/lib/utils";
import { PrefForm } from "@/types/filterTypes";
import { ReactElement } from "react";
import { Control, Controller } from "react-hook-form";

export default function BoolPref({
  control,
  filterTitle,
  keyName,
}: {
  control: Control<PrefForm, unknown>;
  filterTitle: string;
  keyName: string;
}): ReactElement {
  return (
    <Controller
      name={`prefs.${keyName}`}
      control={control}
      defaultValue={false}
      render={({ field: { onChange, onBlur, value, disabled, name, ref } }) => {
        const transformedValue: boolean | undefined =
          typeof value === "boolean" ? value : undefined;
        return (
          <Button
            asChild
            className={cn(
              "bg-secondary flex w-full items-center justify-between gap-2 overflow-hidden border-b border-b-gray-300 p-0",
            )}
          >
            <div className="pr-4">
              <Label
                htmlFor={keyName}
                className="flex h-full w-full cursor-pointer items-center justify-between gap-2 overflow-hidden pl-4"
              >
                <p className="flex grow">{filterTitle}</p>
                <Switch
                  className="border border-gray-400"
                  id={keyName}
                  name={name}
                  ref={ref}
                  checked={transformedValue}
                  onCheckedChange={onChange}
                  onBlur={onBlur}
                  disabled={disabled}
                />
              </Label>
              <PrioritySelector
                key={keyName}
                control={control}
                keyName={keyName}
              />
            </div>
          </Button>
        );
      }}
    />
  );
}
