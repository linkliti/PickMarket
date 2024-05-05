import { z } from "zod";


const PrefFormSchema = z.object({
  params: z.string(),
  userQuery: z.string(),
  priorities: z.record(z.number()),
  prefs: z.record(
    z.union([z.number(), z.array(z.string()), z.boolean()]),
  ),
});

export type PrefForm = z.infer<typeof PrefFormSchema>;
