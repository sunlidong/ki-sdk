/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package instll

// this test mimics the original e2e test with the difference of injecting interface functions implementations
// to programmatically supply configs instead of using a yaml file. With this change, application developers can fetch
// configs from any source as long as they provide their own implementations.
