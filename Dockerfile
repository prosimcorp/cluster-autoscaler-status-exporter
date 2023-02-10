########################################################################################################################
# BASE
########################################################################################################################
FROM debian:buster-slim as base

# Prepare app directory
RUN mkdir -p /usr/app/case/
WORKDIR /usr/app/case/

# Configure entrypoint
COPY ./docker-entrypoint.sh /usr/local/bin/
RUN chmod 0775 /usr/local/bin/docker-entrypoint.sh
ENTRYPOINT ["docker-entrypoint.sh"]
# This step will be replaced by the entrypoint plus the args field defined in the kubernetes eployment manifest
CMD ["sh"]

########################################################################################################################
# BUILD
########################################################################################################################
FROM golang:1.17-buster as build

# Copy the application files
COPY . /usr/app/case/

# Build the application
RUN cd /usr/app/case/ \
    && make build

########################################################################################################################
# APPLICATION
# FOR PROD BUILD ADD TARGET FLAG: docker build . --tag 'case:buster-slim' --target application
########################################################################################################################
FROM base as application

# Copy the build application to the working directory
COPY --from=build /usr/app/case/bin/* /usr/app/case/

# Prepare executable permissions
RUN chmod -R 0775 /usr/app/case/case

# Link application
RUN ln -s /usr/app/case/case /usr/local/bin/case && \
    chmod +x /usr/local/bin/case

########################################################################################################################
# DEBUG
# FOR A DEBUG BUILD: docker build . --tag 'case:buster-slim'
########################################################################################################################
FROM application as debug
# Install debug packages
RUN apt-get update --yes && \
    apt-get install --yes --no-install-recommends \
        bash \
        procps