/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @flow strict-local
 * @format
 */
import type {
  call_trace,
  mutable_call_trace,
} from '../../../generated/MagmaAPIBindings';

import React from 'react';

export type TraceContextType = {
  state: {[string]: call_trace},
  setState?: (key: string, val?: mutable_call_trace) => Promise<void>,
};

export default React.createContext<TraceContextType>({});
