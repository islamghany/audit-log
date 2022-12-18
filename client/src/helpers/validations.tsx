interface Errors {
  [x: string]: any;
}
export const validateEmail = (errors: Errors) =>
  errors?.email
    ? errors.email?.message
      ? errors.email.message
      : errors.email?.type === "required"
      ? "Email can not be empty"
      : "Invalid Email"
    : undefined;

export const validatePassword = (errors: Errors) =>
  errors?.password
    ? errors.password?.message
      ? errors.password.message
      : errors.password?.type === "required"
      ? "Password can not be empty"
      : "Password must be at least 8 and ot be more than 72 chars long"
    : undefined;

export const validateName = (errors: Errors) =>
  errors?.name
    ? errors.name?.message
      ? errors.name.message
      : errors.name?.type === "required"
      ? "Name can not be empty"
      : "Name must be at least 8 and ot be more than 72 chars long"
    : undefined;

export const validateFirstName = (errors: Errors) =>
  errors?.first_name
    ? errors.first_name?.message
      ? errors.first_name.message
      : errors.first_name?.type === "required"
      ? "First Name can not be empty"
      : "Invalid First Name"
    : undefined;

export const validateLastName = (errors: Errors) =>
  errors?.last_name
    ? errors.last_name?.message
      ? errors.last_name.message
      : errors.last_name?.type === "required"
      ? "Last Name can not be empty"
      : "Invalid Last Name"
    : undefined;

export const validateCompanyName = (errors: Errors) =>
  errors?.company_name
    ? errors.compant_name?.message
      ? errors.compant_name.message
      : "Invalid Company Name"
    : undefined;
