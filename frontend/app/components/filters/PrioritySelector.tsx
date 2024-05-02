import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Star } from "lucide-react";

export default function PrioritySelector() {
  return (
    <Select>
      <SelectTrigger className="h-8 w-fit px-2 py-0">
        <SelectValue placeholder={<Star className="h-4 w-4" />} />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectItem value="0">Apple</SelectItem>
          <SelectItem value="1">Apple</SelectItem>
          <SelectItem value="2">Banana</SelectItem>
          <SelectItem value="3">Blueberry</SelectItem>
          <SelectItem value="4">Grapes</SelectItem>
          <SelectItem value="5">Pineapple</SelectItem>
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}
