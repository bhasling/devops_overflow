/*
    RestCall.mjs
    @copyright 2020 William M Hasling All Rights Reserved

    Component to support making a rest call to a remote API rest service.
*/

class RestCall {
    constructor() {
        this.method = null;
        this.url = null;
        this.jsonBody = null;
        this.defaultErrorMessage = "Error occurred";
        this.noAuthorization = false;
        this.simulation = null;
        this.resolve = null;
        this.reject = null;
    }

    invoke(method, url, jsonBody, defaultErrorMessage, simulation) {
        this.method = method;
        this.url = url;
        if (jsonBody) this.jsonBody = jsonBody;
        if (defaultErrorMessage) this.defaultErrorMessage = defaultErrorMessage;
        if (simulation) this.simulation = simulation;
        var result = new Promise((resolve, reject) => {this.execute(resolve, reject)});
        return result;
    }

    execute(resolve, reject) {
        this.resolve = resolve;
        this.reject = reject;
        // If this is the local origin just simulate the response for testing
        if (window && window.location && window.location.origin) {
            var origin = window.location.origin;
            if (origin == "file://" || origin.startsWith("http://localhost")) {
                resolve(this.simulation);
                return;
            }
        }
        // Otherwise, use fetch to execute the remote API call
        try {
            // Create the fetch options
            var options = {
                method: this.method
            };
            if (this.jsonBody != null) {
                options.body = JSON.stringify(this.jsonBody);
            }

            // Invoke fetch, resolve the promises and the returned promise
            fetch(this.url, options)
            .then(response => {
                if (response.status == 200) {
                    this.resolvePromise(response.json());
                } else if (response.status == 400) {
                    this.resolvePromiseWithJsonError(response.json());
                } else {
                    this.resolvePromiseWithError(response.status, response.text());
                }
            });
        } catch (err) {
            this.reject(err);
        }
    }

    resolvePromise(promise) {
        if (this.resolve) {
            // Resolve the fetch promise in the success case
            // Use the result to resolve the promise we returned
            promise.then((response) => {
                this.resolve(response);    
            })
        }
    }

    resolvePromiseWithError(status, promise) {
        // Resolve the fetch promise in the error case
        // Use the result to resolve the reject promise we returned
        if (this.reject) {
            promise.then((data) => {
                var message = data;
                if (!message) message = this.defaultErrorMessage;
                if (!message) message = `Error ${status}.`;
                this.reject(message);
            });
        }
    }
    resolvePromiseWithJsonError(promise) {
        // Resolve the fetch promise in the 400 case
        // Use the json result to resolve the reject promise we returned
        if (this.reject) {
            promise.then((jsonResponse) => {
                var message = jsonResponse.message;
                if (!message) message = this.defaultErrorMessage;
                this.reject(message);    
            })
        }
    }
}

class RestServiceFactory {
    invoke(method, url, jsonBody, defaultErrorMessage, simulation, noAuthorization) {
        // Create a new instance of the RestService object so this method is stateless for parallel threads
        var restService = new RestService();
        var result = restService.invoke(method, url, jsonBody, defaultErrorMessage, simulation, noAuthorization);
        return result;
    }
}
var instance = new RestServiceFactory();
export default instance;
