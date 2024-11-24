FROM golang:1.23.3 AS build

COPY ./ /file-path-validator

WORKDIR /file-path-validator

RUN go mod download && go build -o file_path_validator ./cmd/file_path_validator/


FROM gcr.io/distroless/static-debian12:nonroot-8701094b7fe8ff30d0777bbdfcc9a65caff6f40b

COPY --from=build /file-path-validator/file_path_validator /file_path_validator

HEALTHCHECK --timeout=1s --retries=1 CMD /file_path_validator || exit 1

ENTRYPOINT ["/file_path_validator"]