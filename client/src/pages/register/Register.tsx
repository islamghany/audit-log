import { Button } from "@/components/Button";
import { Input } from "@/components/Form";
import { Title } from "@/components/Typo";
import {
  validateCompanyName,
  validateEmail,
  validateFirstName,
  validateLastName,
  validateName,
  validatePassword,
} from "@/helpers/validations";
import { useForm } from "react-hook-form";

interface RegisterCredintial {
  first_name: string;
  company_name?: string;
  last_name: string;
  email: string;
  password: string;
}
export function Register() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterCredintial>();
  const onSubmit = handleSubmit((e) => {
    console.log(e);
  });
  return (
    <div className="w-full flex items-center mt-20 flex-col">
      <Title
        as="h1"
        className=" mb-6 bg-gradient-to-r from-indigo-200 via-primary to-indigo-200 bg-clip-text font-display text-5xl tracking-tight text-transparent"
      >
        Get started for free
      </Title>
      <div className="max-w-md w-full bg-slate-800 py-4 px-4 rounded-xl">
        <form className="space-y-4" onSubmit={onSubmit}>
          <Input
            name="first_name"
            type="text"
            label="First Name"
            register={register("first_name", {
              required: true,
              pattern: /^[a-z0-9_-]{3,15}$/i,
            })}
            error={validateFirstName(errors)}
          />
          <Input
            name="last_name"
            type="text"
            label="Last Name"
            register={register("last_name", {
              required: true,
              pattern: /^[a-z0-9_-]{3,15}$/i,
            })}
            error={validateLastName(errors)}
          />
          <Input
            name="email"
            type="text"
            label="Email Address"
            register={register("email", {
              required: true,
              pattern: /^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$/i,
            })}
            error={validateEmail(errors)}
          />
          <Input
            name="company_name"
            type="text"
            label="company_name"
            register={register("company_name", {
              minLength: 2,
              maxLength: 72,
            })}
            error={validateCompanyName(errors)}
          />
          <Input
            register={register("password", {
              required: true,
              minLength: 8,
              maxLength: 72,
            })}
            name="password"
            type="password"
            label="Password"
            error={validatePassword(errors)}
          />
          <Button className="w-full">Submit</Button>
        </form>
      </div>
    </div>
  );
}
