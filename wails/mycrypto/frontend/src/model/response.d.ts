type Signature = {
  format: string;
  content: string;
};

type ResponseError = {
  isError: boolean;
  message: string;
};

type Verification = {
  result: string;
  messages: map<string, { result: string }>;
};

interface Response {
  signature: Signature;
  error: ResponseError;
  description?: string;
  verification: Verification;
}

type KeyData = {
  kid: string;
};
